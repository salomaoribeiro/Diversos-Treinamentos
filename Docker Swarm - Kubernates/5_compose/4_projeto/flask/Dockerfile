FROM python:3

RUN apt update -y && apt install -y python3-pip python3-dev

WORKDIR /app

RUN pip install Flask requests flask_mysqldb

COPY . .

EXPOSE 5000

CMD [ "python", "./app.py" ]