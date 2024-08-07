package types

type Cfg_DeliveryChannelSnapshotDeliveryProperties struct {
	// The frequency with which AWS Config recurringly delivers configuration snapshotsE.g., `One_Hour` or `Three_Hours`. Valid values are listed [here](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html#API_ConfigSnapshotDeliveryProperties_Contents).
	DeliveryFrequency string `json:"deliveryFrequency,omitempty" yaml:"deliveryFrequency,omitempty"`
}
