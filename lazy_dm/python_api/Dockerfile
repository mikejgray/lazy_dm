FROM public.ecr.aws/lambda/python:3.8

COPY ./lazy_dm_api ./app
COPY ./requirements.txt ./requirements.txt
RUN pip install -r ./requirements.txt

CMD ["app.main.handler"]