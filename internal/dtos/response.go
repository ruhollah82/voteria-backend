package dtos

type ResponseDTO struct {
	UserErrs     []error
	ServerErr    error
	Data         any // struct or nil
	ResponseCode string
	Status       int
	Msg string
}
