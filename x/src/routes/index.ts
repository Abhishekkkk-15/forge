import { Router } from "express";
import healthRoute from "./health.route";

const router = Router();

router.use("/health", healthRoute);

export default router;
