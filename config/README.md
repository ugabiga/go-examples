# Thoughts on Configuration

## 1. YAML (gopkg)

If you’re planning to use only YAML for your configuration, the easiest route is to go with the gopkg.in/yaml package.
It’s straightforward and user-friendly, making it a breeze to work with.

## 2. Viper

If you need to handle multiple configuration formats, spf13/viper is your best bet.
Not only does it support various formats, but it also handles environment variables,
giving you a lot of flexibility in how you manage your configurations.
But getting environment variables to work with Viper is a bit tricky.

