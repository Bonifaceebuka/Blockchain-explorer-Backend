package helpers

func GetMessage(message_type string) string {

	switch message_type {
	case "fetch_latest_block":
		return "Unable to fetch the latest block information!"
	case "fetch_a_block":
		return "Unable to convert Block ID:"
	default:
		return ""
	}
}
