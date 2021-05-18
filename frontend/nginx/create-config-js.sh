cat >/app/config.js <<EOL
const config = (() => {
  return {
    "API_ENDPOINT": "${API_ENDPOINT}",
  }
})();
EOL