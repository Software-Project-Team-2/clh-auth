# Go Authentication Service

This project provides a basic authentication service using JWT (JSON Web Tokens) for secure token generation and validation, and Redis for user management. It includes functionalities for user creation, token generation, and token validation, with support for gRPC and Protocol Buffers.

## Installation

Before installing this project, ensure you have Docker installed on your system.

1. Clone the repository:
   ```bash
   git clone [URL of your repository]
   ```

2. Navigate to the project directory:
   ```bash
   cd [project directory]
   ```

## Building the Docker Container

To build the Docker container for this project, run:

```bash
docker build -t clh-auth .
```

Replace `clh-auth` with your desired application name.

## Running the Docker Container

To run the Docker container, use the following command:

```bash
docker run -d --name clh-auth -e JWT_SECRET_TOKEN=your_jwt_secret -e IS_PRODUCTION=false -e REDIS_HOST=localhost:6379 -e REDIS_PASSWORD= -e REDIS_DB=0 -p 8080:8080 glad2os/clh-auth
```

### Environment Variables

- `JWT_SECRET_TOKEN`: Secret token for JWT. Default is `ahd8fee2ohboTh8eS9eeyoosaine3ohK` if not set.
- `IS_PRODUCTION`: Set to `true` for production environments. Ensure `JWT_SECRET_TOKEN` is set in production.
- `REDIS_HOST`: The host address of the Redis server. Required.
- `REDIS_PASSWORD`: Password for Redis server. Default is empty.
- `REDIS_DB`: Redis database number. Default is `"0"`.

## Contributing

Contributions are welcome. Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature-branch-name`.
3. Make your changes and commit them: `git commit -m 'commit message'`.
4. Push to the original branch: `git push origin [project name]/[location]`.
5. Create the pull request.

Alternatively, see the GitHub documentation on [creating a pull request](https://help.github.com/articles/creating-a-pull-request/).

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.
