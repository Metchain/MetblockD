package external

// NewBlockGHOSTDAGData creates a new instance of BlockGHOSTDAGData
func NewMETGBlockData(

	selectedParent *DomainHash,
) *MetGDagData {

	return &MetGDagData{

		selectedParent: selectedParent,
	}
}

// SelectedParent returns the SelectedParent of the block
func (mgdb *MetGDagData) SelectedParent() *DomainHash {
	return mgdb.selectedParent
}
