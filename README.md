# Discord Lookup

It is inspired by the api of mesalytic [discord-lookup-api](https://github.com/mesalytic/discord-lookup-api).

This API allows you to retrieve detailed information about Discord users, applications, and guilds using their unique IDs.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Environment Variables](#environment-variables)
- [Installation](#installation)
- [Running the Server](#running-the-server)
- [Contributing](#contributing)
- [License](#license)

## Features

- **User Information**: Retrieve user ID, tag, badges, avatar, and banner details.
- **Application Information**: Get application ID, name, icon, description, and more.
- **Guild Information**: Fetch guild ID, name, instant invite, and presence count.
- **CORS Support**: Built-in support for cross-origin requests.
- **Caching**: Data is cached for 3 hours to improve performance.

## Getting Started

To start using the Discord Lookup API, you can deploy your own.

## Installation

Follow these steps to set up the project locally:

1. **Clone the repository**:

   ```sh
   git clone https://github.com/polynux/discord-lookup.git
   cd discord-lookup
   ```

2. **Install dependencies**:

   ```sh
   npm install
   ```

3. **Set up environment variables**:
   Create a `.env` file in the root directory using `cp .env.example .env` and adjust variables accordingly.

4. **Build the binary**

   `./build.sh`

## Running the Server

Start the server using the following command:

```sh
./build/discord-lookup
```

The server will start on the port specified in the `.env.local` file (default is 3000).

## API Endpoints

### User Lookup

- **Endpoint**: `/api/v1/user/:id`
- **Method**: GET
- **Description**: Retrieve information about a Discord user by their ID.
- **Example**:
  ```sh
  curl https:/example.com/api/v1/user/123456789012345678
  ```
  ```json
  {
      "id": "123456789012345678",
      "tag": "example#0001",
      "badges": ["HOUSE_BRAVERY", "EARLY_VERIFIED_BOT_DEVELOPER"],
      "avatar": {
          "id": "avatar_id",
          "link": "https://cdn.discordapp.com/avatars/123456789012345678/avatar_id",
          "is_animated": false
      },
      "banner": {
          "id": "banner_id",
          "link": "https://cdn.discordapp.com/banners/123456789012345678/banner_id",
          "is_animated": true,
          "color": "#123456"
      }
  }
  ```

### Application Lookup

- **Endpoint**: `/api/v1/application/:id`
- **Method**: GET
- **Description**: Get details about a Discord application using its ID.
- **Example**:
  ```sh
  curl https://example.com/api/v1/application/123456789012345678
  ```
  ```json
  {
      "id": "123456789012345678",
      "name": "ExampleApp",
      "icon": "https://cdn.discordapp.com/avatars/123456789012345678/icon_id",
      "description": "This is an example application.",
      "summary": "",
      "type": null,
      "hook": true,
      "guild_id": "123456789012345678",
      "bot_public": true,
      "bot_require_code_grant": false,
      "terms_of_service_url": "https://example.com/tos",
      "privacy_policy_url": "https://example.com/privacy",
      "install_params": {
          "scopes": ["bot", "applications.commands"],
          "permissions": "1234567890"
      },
      "verify_key": "verify_key",
      "flags": {
          "bits": 12345678,
          "detailed": ["GATEWAY_GUILD_MEMBERS", "GATEWAY_MESSAGE_CONTENT"]
      },
      "tags": ["example", "bot"]
  }
  ```

### Guild Lookup

> **Note**: The guild must have Server Widget and/or Server Discovery enabled.

- **Endpoint**: `/api/v1/guild/:id`
- **Method**: GET
- **Description**: Fetch information about a Discord guild by its ID.
- **Example**:
  ```sh
  curl https://example.com/api/v1/guild/123456789012345678
  ```

  ```json
  {
      "id": "123456789012345678",
      "name": "Example Guild",
      "instant_invite": "https://discord.gg/example",
      "presence_count": 1234
  }
  ```

## Contributing

We welcome contributions! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## License

This project is licensed under the GPL V3. See the [LICENSE](LICENSE) file for details.

