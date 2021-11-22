package python

var dockerfileText = `
FROM python:3.9 as builder

ARG GIT_REFERENCE
ARG SSH_KEY
ENV GIT_REFERENCE=$GIT_REFERENCE
ENV PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1 

RUN apt-get update \
    && apt-get install openssh-client git build-essential -y && rm -rfv /var/cache/apt/* \
    && mkdir ~/.ssh \
    && ssh-keyscan github.com > /root/.ssh/known_hosts \
    && eval $(ssh-agent) \
    && echo "${SSH_KEY}" | ssh-add - \
    && apt-get install curl 

RUN pip3 install poetry
COPY {{.appWithUnderScore}} /{{.appWithUnderScore}}
COPY poetry.lock pyproject.toml ./
RUN poetry install --no-dev
RUN poetry build --format wheel

FROM python:3.9-slim
COPY --from=builder /dist/* /dist/
RUN ls /dist
RUN bash -c "pip3 install /dist/*.whl"

EXPOSE 5000
`
