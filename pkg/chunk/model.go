package chunk

const (
	MetadataPartitionKey = "Partition"
)

type Config struct {
	MaxSliceSize int `json:"maxSliceSize"`
	Partitions   int `json:"partitions"`
}
