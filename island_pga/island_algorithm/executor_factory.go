package island_algorithm

type SimpleIslandExecutorFactory struct {
	CriteriaFactory OptimizationCriteriaFactory
}

func (factory *SimpleIslandExecutorFactory)CreateExecutor(criterionType OptimizationCriterionType, recipeIds []int) (IslandExecutor, error) {
	criterion := factory.CriteriaFactory.CreateCriterion(criterionType, recipeIds)
	executor := NewSimpleIslandExecutorImpl(criterion)
	return executor, nil
}