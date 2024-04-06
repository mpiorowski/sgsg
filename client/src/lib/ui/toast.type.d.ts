export type Toast = {
    id: string;
    type: "success" | "error" | "warning" | "info";
    title: string;
    description?: string;
    duration: number;
    action?: {
        label: string;
        onClick: () => void;
    };
};
