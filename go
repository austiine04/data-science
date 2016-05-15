!#bin/bash

docker-machine start pneuma
eval "$(docker-machine env pneuma)"
docker run -p 8888:8888 -v /Users/kaugust/GADataScience/SYD_DAT_4/project:/tmp/project data
