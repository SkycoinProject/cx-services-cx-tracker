FROM node:10.15.0-alpine AS build

# Set work directory
WORKDIR /usr/src/cx-tracker-web/
# Add frontend code
ADD web /usr/src/cx-tracker-web/

RUN npm install -g @angular/cli \
    && npm install \
    && ng build --prod --build-optimizer --vendor-chunk=true

FROM nginx:alpine

COPY --from=build /usr/src/cx-tracker-web/dist/ /usr/share/nginx/html/dist

COPY ./docker/images/cx-tracker-web/nginx.conf /etc/nginx/nginx.conf
