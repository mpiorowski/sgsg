import { z } from "zod";

export const schema = z.object({
    id: z.string().min(1, "required"),
    title: z.string().min(1, "required"),
    content: z.string().min(1, "required"),
});
