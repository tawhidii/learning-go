package apis

type UserAPIType struct {
	ApiVersion float64
	ApiDetails string
}

func UserAPIDetails(apiDetails string) string {
	return apiDetails
}
