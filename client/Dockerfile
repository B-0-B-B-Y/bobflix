FROM node:alpine as dependencies
RUN apk add --no-cache libc6-compat
WORKDIR /client
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

FROM node:alpine as builder
WORKDIR /client
COPY . .
COPY --from=dependencies /client/node_modules ./node_modules
RUN yarn build && yarn install --production --ignore-scripts --prefer-offline

FROM node:alpine as runner
WORKDIR /client
ENV NODE_ENV production

RUN addgroup -g 1001 -S nodejs
RUN adduser -S nextjs -u 1001

COPY --from=builder /client/next.config.js ./
COPY --from=builder /client/public ./public
COPY --from=builder /client/.next ./.next
COPY --from=builder /client/node_modules ./node_modules
COPY --from=builder /client/package.json ./package.json

USER nextjs
ENV NEXT_TELEMETRY_DISABLED 1

EXPOSE 3000
CMD ["yarn", "start"]