package island_algorithm


type SimpleOptimizationCriteriaFactoryImpl struct {
	InfoExtractor RecipeInfoExtractor
}

func (factory *SimpleOptimizationCriteriaFactoryImpl)CreateCriterion(criterionType OptimizationCriterionType) CriterionEvaluator {
	switch criterionType {
	case MinPrice:
		return &MinPriceCriterionEvaluator{ recipeInfoExtractor: factory.InfoExtractor }
	case MinCookingTime:
		return &MinCookingTimeCriterionEvaluator{ recipeInfoExtractor: factory.InfoExtractor }
	case MaxProteins:
		return &MaxProteinsCriterionEvaluator{ recipeInfoExtractor: factory.InfoExtractor }
	default:
		return &MinPriceCriterionEvaluator{ recipeInfoExtractor: factory.InfoExtractor } // default case to silence compiler
	}
}