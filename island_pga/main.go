package main

import (
	"fmt"
	"log"
	"database/sql"
	"net/http"
	"cookhub.com/app/db"

	"island_algorithm"
	"ga_operators"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type OptimizationResponse struct {
	RecipeIds [][]int `json:"recipe_ids"`
}

func main() {
	
	var database *sql.DB
	database, err := db.InitStore()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	recipeInfoExtractor := island_algorithm.RecipeInfoExtractor {
		Database: database,
	}
	criterionFactory := island_algorithm.SimpleOptimizationCriteriaFactoryImpl {
		InfoExtractor: recipeInfoExtractor,
	}
	executorFactory := island_algorithm.SimpleIslandExecutorFactory { 
		CriteriaFactory: &criterionFactory,
	}

	islandMaster := island_algorithm.SimpleIslandMasterImpl {
		ExecutorFactory: &executorFactory,
	}

	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.GET("/", func (context echo.Context) error {
		return context.HTML(http.StatusOK, fmt.Sprintf("Hello, CookHub optimizer!"))
	})

	server.GET("/optimize", func (context echo.Context) error {
		log.Print("GA_OPTIMIZER::Started execution")
		recipeIds, requirements, err := recipeInfoExtractor.GetChromosomeRequirements()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to read requirements using database: %s", err))
		}
		
		resultChannel := islandMaster.StartOptimization(island_algorithm.GeneticAlgorithmParams {
			TotalIterationsCount: 200,
			MigrationIterationsCount: 50,
			Operators: []island_algorithm.GeneticOperatorMetadata {
				island_algorithm.GeneticOperatorMetadata {
					Type: ga_operators.Inversion,
					Probability: 0.3,
				},
				island_algorithm.GeneticOperatorMetadata {
					Type: ga_operators.Mutation,
					Probability: 0.4,
				},
				island_algorithm.GeneticOperatorMetadata {
					Type: ga_operators.SinglePointCrossingover,
					Probability: 0.3,
				},
			},
			Requirements: *requirements,
			RecipesIds: recipeIds,
			IslandsCount: 3,
		})

		optimizationRes := <-resultChannel
		var bestSolutions [][]int
		for i:=0; i < len(optimizationRes.BestSolutions); i++ {
			curSolution := make([]int, 0)
			chromosomeValues := optimizationRes.BestSolutions[i].GetValues()
			for j:=0; j < len(chromosomeValues); j++ {
				if chromosomeValues[j] == 1 {
					curSolution = append(curSolution, recipeIds[j])
				}
			}
			bestSolutions = append(bestSolutions, curSolution)
		}

		response := OptimizationResponse {
			RecipeIds: bestSolutions,
		}

		return context.JSON(http.StatusOK, response)
	})

	server.Logger.Fatal(server.Start(":8090"))
}
