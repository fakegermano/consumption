# FROM tiangolo/node-frontend:14 as build-stage

# WORKDIR /app

# COPY package*.json /app/

# RUN npm install

# COPY ./meterusage /app/

# RUN npm run build

# FROM nginx:latest
# MAINTAINER Daniel Travieso <danielgtravieso@gmail.com>

# COPY --from=build-stage /app/build /usr/share/nginx/html
# COPY --from=build-stage /nginx.conf /etc/nginx/conf.d/default.conf


FROM busybox

CMD ["sleep", "60"]
