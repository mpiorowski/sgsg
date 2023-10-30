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

export type UpsendFile = {
    id: string;
    created: string;
    updated: string;
    deleted: string;
    user_id: string;
    name: string;
    size: string;
    mime_type: string;
    buffer: string[];
    base64: string;
};

export type UpsendImage = {
    id: string;
    created: string;
    updated: string;
    deleted: string;
    user_id: string;
    name: string;
    size: string;
    mime_type: string;
    target_id: string;
    url: string;
    buffer: string[];
};
