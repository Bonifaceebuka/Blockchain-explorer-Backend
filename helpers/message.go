package helpers

func GetMessage(message_type string) string {
	message := ""
	switch message_type {
	case "cant_convert_value":
		message = "Unable to convert the selected value!"
	case "fetch_a_block":
		message = "Unable to convert Block ID:"
	case "tnx_is_pending":
		message = "Transaction is pending"
	case "invalid_tnx_hash":
		message = "Invalid Transaction hash"
	}

	return message
}
