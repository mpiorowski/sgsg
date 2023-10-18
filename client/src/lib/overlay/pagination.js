/** Generate pagination schema
 * @template T
 * @param {T[]} data
 * @param {number} currentPage
 * @param {number} pageSize
 * @returns {number[]}
 */
function generatePaginationSchema(data, currentPage, pageSize = 10) {
    const totalItems = data.length;
    const totalPages = Math.ceil(totalItems / pageSize);
    const paginationArray = [];

    if (totalPages <= 7) {
        for (let i = 1; i <= totalPages; i++) {
            paginationArray.push(i);
        }
    } else {
        paginationArray.push(1);

        if (currentPage < 4) {
            for (let i = 2; i <= 4; i++) {
                paginationArray.push(i);
            }
            paginationArray.push(0);
        }

        if (currentPage >= 4 && currentPage <= totalPages - 3) {
            paginationArray.push(0);
            for (let i = currentPage - 1; i <= currentPage + 1; i++) {
                paginationArray.push(i);
            }
            paginationArray.push(0);
        }

        if (currentPage > totalPages - 3) {
            paginationArray.push(0);
            for (let i = totalPages - 3; i <= totalPages - 1; i++) {
                paginationArray.push(i);
            }
        }

        paginationArray.push(totalPages);
    }

    return paginationArray;
}

/**
 * Generate pagination data
 * @template T
 * @param {T[]} data
 * @param {number} currentPage
 * @param {number} pageSize
 * @returns {{
 * data: T[],
 * start: number,
 * end: number,
 * prev: number,
 * next: number,
 * total: number,
 * schema: number[]
 * }}
 */
export function pagination(data, currentPage, pageSize = 10) {
    const paginated = data.slice(
        (currentPage - 1) * pageSize,
        currentPage * pageSize,
    );
    const total = data.length;
    const start = (currentPage - 1) * pageSize + 1;
    const end = currentPage * pageSize > total ? total : currentPage * pageSize;
    const prev = currentPage > 1 ? currentPage - 1 : 1;
    const next =
        currentPage < Math.ceil(total / pageSize)
            ? currentPage + 1
            : Math.ceil(total / pageSize);
    const schema = generatePaginationSchema(data, currentPage);

    return {
        data: paginated,
        start,
        end,
        prev,
        next,
        total,
        schema,
    };
}
