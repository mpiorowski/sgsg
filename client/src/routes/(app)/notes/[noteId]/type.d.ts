import { z } from "zod";
import { schema } from "./note";

export type Note = z.infer<typeof schema>;
