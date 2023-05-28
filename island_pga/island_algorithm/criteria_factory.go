package island_algorithm


type SimpleOptimizationCriteriaFactoryImpl struct {
	InfoExtractor RecipeInfoExtractor
}

func (factory *SimpleOptimizationCriteriaFactoryImpl)CreateCriterion(criterionType OptimizationCriterionType, recipeIds []int) CriterionEvaluator {
	switch criterionType {
	case MinPrice:
		return &MinPriceCriterionEvaluator{ recipeInfoExtractor: factory.InfoExtractor, recipeIds: recipeIds }
	case MinCookingTime:
		return &MinCookingTimeCriterionEvaluator{ recipeInfoExtractor: factory.InfoExtractor, recipeIds: recipeIds }
	case MaxProteins:
		return &MaxProteinsCriterionEvaluator{ recipeInfoExtractor: factory.InfoExtractor, recipeIds: recipeIds }
	default:
		return &MinPriceCriterionEvaluator{ recipeInfoExtractor: factory.InfoExtractor, recipeIds: recipeIds } // default case to silence compiler
	}
}