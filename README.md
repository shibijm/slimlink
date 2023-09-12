# Slimlink

Slimlink is a simple and lightweight web service for shortening URLs.

[![Latest Release](https://img.shields.io/github/v/release/shibijm/slimlink?label=Latest%20Release)](https://github.com/shibijm/slimlink/releases/latest)
[![Build Status](https://img.shields.io/github/actions/workflow/status/shibijm/slimlink/release.yml?label=Build&logo=github)](https://github.com/shibijm/slimlink/actions/workflows/release.yml)
[![Deployment Status](https://img.shields.io/website/https/l.shjm.in?label=Deployment&logo=docker)](https://l.shjm.in)

## Usage

A public demo instance is hosted at [l.shjm.in](https://l.shjm.in).

To host your own instance, you can use [Docker](https://github.com/shibijm/slimlink/pkgs/container/slimlink) or download a build from the [releases page](https://github.com/shibijm/slimlink/releases).

## Configuration - Environment Variables

* `BIND_ADDRESS` (default: `0.0.0.0`)
* `BIND_PORT` (default: `44558`)
* `LINK_ID_LENGTH` (default: `5`, min: `1`, max: `64`) - Length of each random Base62 string which gets generated for use in shortened URLs. 916 million (62<sup>5</sup>) possibilities when set to 5.
* `REDIS_CONNECTION_STRING` (format: `redis://USERNAME:PASSWORD@HOST:PORT`)
* `MYSQL_CONNECTION_STRING` (format: `USERNAME:PASSWORD@tcp(HOST:PORT)/DATABASE_NAME`)
* `PAGE_FOOTER_TEXT` (optional) - Text to be displayed at the bottom-left corner of the main page.

Environment variables will be automatically loaded from a `.env` file if one exists in the program's working directory.

Only one of the database connection strings must be set.
