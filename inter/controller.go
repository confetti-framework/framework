package inter

type PipeHolder func(request Request) Response
type Controller = PipeHolder
