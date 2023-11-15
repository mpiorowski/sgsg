import { z } from "zod";

export type Note = z.infer<typeof schema>;
