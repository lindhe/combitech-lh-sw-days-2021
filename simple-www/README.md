# Simple WWW

I'm _not_ a web designer.
Red, green or blue is the best I can give you.

## Build

Using `podman` instead of `docker` is cool.
You want to be cool.
They use pretty much the same syntax, just slightly different defaults from a user's perspective.

### Using Podman

```console
$ podman build -t docker.io/lindhe/simple-www:red-0.1.0 .
STEP 1/2: FROM docker.io/nginx:1.21.3-alpine
STEP 2/2: COPY www/* /usr/share/nginx/html
--> Using cache 47e4e7b115091dc22fd0ba49885c026df3721102ee61a073acf029edf90f35c7
COMMIT docker.io/lindhe/simple-www:red-0.1.0
--> 47e4e7b1150
Successfully tagged docker.io/lindhe/simple-www:red-0.1.0
Successfully tagged localhost/www:latest
47e4e7b115091dc22fd0ba49885c026df3721102ee61a073acf029edf90f35c7
```

### Using Docker

```console
$ docker build -f Containerfile -t docker.io/lindhe/simple-www:red-0.1.0 .
Sending build context to Docker daemon  6.656kB
Step 1/2 : FROM docker.io/nginx:1.21.3-alpine
1.21.3-alpine: Pulling from library/nginx
a0d0a0d46f8b: Pull complete
4dd4efe90939: Pull complete
c1368e94e1ec: Pull complete
3e72c40d0ff4: Pull complete
969825a5ca61: Pull complete
61074acc7dd2: Pull complete
Digest: sha256:1ff1364a1c4332341fc0a854820f1d50e90e11bb0b93eb53b47dc5e10c680116
Status: Downloaded newer image for nginx:1.21.3-alpine
 ---> 513f9a9d8748
Step 2/2 : COPY www/* /usr/share/nginx/html
 ---> 125abc7b1f83
Successfully built 125abc7b1f83
Successfully tagged lindhe/simple-www:red-0.1.0
```

## Push

```console
$  echo "${DOCKER_PASSWORD}" | podman login -u "${DOCKER_USERNAME}" --password-stdin docker.io
Login Succeeded!
$ podman push docker.io/lindhe/simple-www:red-0.1.0
Getting image source signatures
Copying blob e2eb06d8af82 skipped: already exists
Copying blob e6d3cea19fef skipped: already exists
Copying blob 40403bebe4fd skipped: already exists
Copying blob 20d0effdf3a2 skipped: already exists
Copying blob 311d8db33235 skipped: already exists
Copying blob b4b4e85910ea skipped: already exists
Copying blob 38d4352aecf2 done
Copying config 47e4e7b115 done
Writing manifest to image destination
Storing signatures
```
