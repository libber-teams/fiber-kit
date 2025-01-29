package errors

type HttpError struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`

	wrappedError error
} // @name HttpError

func NewHttpError(
	typ, title, detail, instance string,
	status int,
) *HttpError {
	return &HttpError{
		Type:     typ,
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
	}
}

func NewHttpNotFoundError(
	instance string,
) *HttpError {
	return NewHttpError(
		"not_found",
		"Recurso não encontrado",
		"O recurso solicitado não foi encontrado.",
		instance,
		404,
	)
}

func NewHttpForbiddenError(
	instance string,
) *HttpError {
	return NewHttpError(
		"forbidden",
		"Acesso negado",
		"Você não tem permissão para acessar este recurso.",
		instance,
		403,
	)
}

func NewHttpBadRequestError(
	instance string,
) *HttpError {
	return NewHttpError(
		"bad_request",
		"Requisição inválida",
		"A requisição não pode ser processada devido a erros de validação.",
		instance,
		400,
	)
}

func NewHttpInternalServerError(
	instance string,
) *HttpError {
	return NewHttpError(
		"internal_server_error",
		"Erro interno do servidor",
		"O servidor encontrou um erro inesperado e não pode processar a requisição.",
		instance,
		500,
	)
}

func NewConflictError(
	instance string,
) *HttpError {
	return NewHttpError(
		"conflict",
		"Conflito",
		"O recurso solicitado já existe.",
		instance,
		409,
	)
}

func NewUnprocessableEntityError(
	instance string,
) *HttpError {
	return NewHttpError(
		"unprocessable_entity",
		"Entidade não processável",
		"A requisição não pode ser processada devido a erros de validação.",
		instance,
		422,
	)
}

func (dto *HttpError) Error() string {
	return dto.Detail
}

func (dto *HttpError) SetType(typ string) *HttpError {
	dto.Type = typ
	return dto
}

func (dto *HttpError) SetTitle(title string) *HttpError {
	dto.Title = title
	return dto
}

func (dto *HttpError) SetStatus(status int) *HttpError {
	dto.Status = status
	return dto
}

func (dto *HttpError) SetDetail(detail string) *HttpError {
	dto.Detail = detail
	return dto
}

func (dto *HttpError) SetInstance(instance string) *HttpError {
	dto.Instance = instance
	return dto
}

func (dto *HttpError) Wrap(err error) *HttpError {
	dto.wrappedError = err
	return dto
}
