# Slimlink

Slimlink is a simple and lightweight web service for shortening URLs.

[![Latest Release](https://img.shields.io/github/v/release/shibijm/slimlink?label=Latest%20Release)](https://github.com/shibijm/slimlink/releases/latest)
[![Build Status](https://img.shields.io/github/actions/workflow/status/shibijm/slimlink/build-and-release.yml?label=Build&logo=github)](https://github.com/shibijm/slimlink/actions/workflows/build-and-release.yml)
[![Deployment Status](https://img.shields.io/github/deployments/shibijm/slimlink/Production?label=Deployment&logo=vercel)](https://github.com/shibijm/slimlink/deployments?environment=Production)

## Usage

A public instance is hosted at [slimlink.vercel.app](https://slimlink.vercel.app).

You can download a build from the [releases](https://github.com/shibijm/slimlink/releases) page and host your own instance.

## Configuration - Environment Variables

* `BIND_ADDRESS` (default: `127.0.0.1`)
* `BIND_PORT` (default: `44558`)
* `LINK_ID_LENGTH` (default: `5`, min: `1`, max: `64`) - Length of each random Base62 string which gets generated for use in shortened URLs. 916 million (62<sup>5</sup>) possibilities when set to 5.
* `REDIS_CONNECTION_STRING` (format: `redis://USERNAME:PASSWORD@HOST:PORT`)
* `MYSQL_CONNECTION_STRING` (format: `USERNAME:PASSWORD@tcp(HOST:PORT)/DATABASE_NAME`)

Environment variables will be automatically loaded from a `.env` file if one exists in the program's working directory.

Only one of the database connection strings must be set.
