#!/bin/bash

image=ascend6/httpenv

docker build -t ${image} .
docker login -u ascend6
docker push ${image}

