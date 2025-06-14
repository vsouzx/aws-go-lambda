FROM golang:1.22 as build

WORKDIR /app
COPY . .

RUN go build -o main

FROM public.ecr.aws/lambda/go:1
COPY --from=build /app/main ${LAMBDA_TASK_ROOT}

CMD ["main"]