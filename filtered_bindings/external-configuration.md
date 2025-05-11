---
id: external-configuration
last_modified: "2025-05-03"
derived_from: no-secret-suppression
enforced_by: linters & secret scanning
applies_to:
  - all
---

# Binding: Externalize All Configuration

Configuration values that vary between environments or are sensitive MUST be externalized and NEVER hardcoded in source code. This includes database strings, API keys/endpoints, ports, feature flags, and any deployment-specific settings.

## Rationale

Hardcoded configuration creates several problems: it ties the code to specific environments, exposes sensitive information in version control, requires source code changes for configuration updates, and complicates deployment to multiple environments. Externalizing configuration follows the principle of separation of concerns by keeping environment-specific settings separate from application logic.

## Enforcement

This binding is enforced by:

1. Secret scanning in pre-commit hooks and CI
2. Linter rules that detect hardcoded configuration patterns
3. Code review processes that reject hardcoded values
4. Architecture design that requires external configuration injection

## Implementation

1. **Use Environment Variables** for deployment flexibility:
   - Primary method for production environments
   - Support for cloud-native deployment models

2. **Use Configuration Files** for local development:
   - `.env` files (with .env.example checked into source control)
   - YAML/JSON configuration files
   - Make sure to exclude actual config files with sensitive values from version control

3. **Load via Libraries** into strongly-typed objects:
   - TypeScript/JavaScript: dotenv, convict
   - Go: viper, godotenv
   - Other languages: similar configuration management libraries

4. **Validate Configuration**:
   - Validate configuration at startup
   - Fail fast if required configuration is missing
   - Provide clear error messages about what's missing or invalid

## Examples

```typescript
// ❌ BAD: Hardcoded configuration
const dbConnection = "postgresql://user:password@localhost:5432/mydb";
const apiKey = "1a2b3c4d5e6f";

// ✅ GOOD: Configuration from environment variables
// config.ts
interface Config {
  database: {
    url: string;
    poolSize: number;
  };
  api: {
    key: string;
    url: string;
  };
}

export const config: Config = {
  database: {
    url: process.env.DATABASE_URL || "postgresql://user:password@localhost:5432/mydb_dev",
    poolSize: parseInt(process.env.DB_POOL_SIZE || "10", 10)
  },
  api: {
    key: process.env.API_KEY || "",
    url: process.env.API_URL || "https://api.dev.example.com"
  }
};

// Validate required configuration
if (!config.api.key) {
  throw new Error("API_KEY environment variable is required");
}
```

```go
// ❌ BAD: Hardcoded configuration
func Connect() (*sql.DB, error) {
    return sql.Open("postgres", "postgresql://user:password@localhost:5432/mydb")
}

// ✅ GOOD: Configuration from environment
// config/config.go
type Config struct {
    Database struct {
        URL      string `mapstructure:"url"`
        PoolSize int    `mapstructure:"pool_size"`
    }
    API struct {
        Key string `mapstructure:"key"`
        URL string `mapstructure:"url"`
    }
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config") // config.yaml
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()
    
    // Map environment variables
    viper.BindEnv("database.url", "DATABASE_URL")
    viper.BindEnv("database.pool_size", "DB_POOL_SIZE")
    viper.BindEnv("api.key", "API_KEY")
    viper.BindEnv("api.url", "API_URL")
    
    // Set defaults
    viper.SetDefault("database.pool_size", 10)
    viper.SetDefault("api.url", "https://api.dev.example.com")
    
    // Read config
    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return nil, err
        }
        // Config file not found - using defaults and env vars only
    }
    
    // Parse config
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }
    
    // Validate required config
    if config.API.Key == "" {
        return nil, errors.New("API key is required")
    }
    
    return &config, nil
}
```

## Secrets Management

For production environments, prefer dedicated secrets management systems:

- Cloud provider solutions (AWS Secrets Manager, GCP Secret Manager, Azure Key Vault)
- HashiCorp Vault
- Kubernetes Secrets (with proper encryption at rest)

## Related Bindings

- [no-secret-suppression.md](./no-secret-suppression.md) - Similar principles for secret management
- [hex-domain-purity.md](./hex-domain-purity.md) - Configuration is an infrastructure concern