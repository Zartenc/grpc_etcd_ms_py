FROM python:3.6
ENV PATH /usr/local/bin:$PATH
COPY . /code
WORKDIR /code
RUN pip install -r requirements.txt
EXPOSE 56789
ENTRYPOINT ["python3", "-u", "main.py"]
