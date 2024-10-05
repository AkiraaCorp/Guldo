![tumblr_inline_pq138vKa7y1r15usg_1280](https://github.com/user-attachments/assets/2ca43911-585a-47ad-9187-95ec577c060c)

Guldo is a Go-based application that connects to a blockchain and a PostgreSQL database to fetch and store event probabilities. It utilizes several packages for database interaction, blockchain communication, and utility functions.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Environment Variables](#environment-variables)
- [Dependencies](#dependencies)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/guldo.git
   cd guldo
   ```

2. **Install dependencies:**

   Ensure you have Go installed, then run:

   ```bash
   go mod tidy
   ```

3. **Set up the environment variables:**

   Create a `.env` file in the root directory and add the necessary environment variables (see [Environment Variables](#environment-variables) section).

## Usage

Run the application using:

## Project Structure

- **`main.go`**: The entry point of the application.
- **`utils/banner.go`**: Contains utility functions, such as printing the application banner.
- **`db/db.go`**: Handles database connections using GORM.
- **`blockchain/blockchain.go`**: Manages blockchain client creation and interaction.
- **`repository/odds_repository.go`**: Manages CRUD operations for odds data.
- **`models/odds.go`**: Defines the data model for odds history.

## Environment Variables

The application requires the following environment variables:

- `DB_HOST`: Database host.
- `DB_PORT`: Database port.
- `DB_USER`: Database user.
- `DB_PASSWORD`: Database password.
- `DB_NAME`: Database name.
- `DB_SSLMODE`: SSL mode for database connection.
- `RPC_URL`: URL for the blockchain RPC.

## Dependencies

The project uses the following Go packages:

- `github.com/NethermindEth/starknet.go`: For blockchain interaction.
- `github.com/joho/godotenv`: For loading environment variables.
- `gorm.io/driver/postgres` and `gorm.io/gorm`: For database interaction.

For a complete list of dependencies, see the `go.mod` file.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
