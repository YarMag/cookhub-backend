package island_algorithm

type SimpleIslandExecutorFactory struct {
	CriteriaFactory OptimizationCriteriaFactory
}

func (factory *SimpleIslandExecutorFactory)CreateExecutor(criterionType OptimizationCriterionType) (IslandExecutor, error) {
	criterion := factory.CriteriaFactory.CreateCriterion(criterionType)
	executor := NewSimpleIslandExecutorImpl(criterion)
	return executor, nil
}