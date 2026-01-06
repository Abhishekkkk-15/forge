import { createApp } from "./app";
import { env } from "./config/env";

const app = createApp();

app.listen(env.PORT, () => {
  console.log(
    `[${env.NODE_ENV}] {{.ProjectName}} running on port ${env.PORT}`
  );
});
