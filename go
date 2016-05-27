#!/bin/bash

docker run -v $(pwd):/home/ds/notebooks -p 8888:8888 dataquestio/python3-starter
