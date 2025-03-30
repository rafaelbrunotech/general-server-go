namespace Constants {
  enum Environment {
    DEV: "development",
    STG: "staging",
    PROD: "production"
  }

  enum Versions {
    v1: "1.0"
  }

  alias LocalhostServer = "http://localhost:8080";
  alias DevelopmentServer = "https://dev.example.com";
  alias StagingServer = "https://stg.example.com";
  alias ProductionServer = "https://example.com";

  alias ContentTypes = "application/json";
}