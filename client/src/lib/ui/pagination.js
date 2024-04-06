/** Generate pagination schema
 * @param {number} total
 * @param {number} page
 * @param {number} limit
 * @returns {number[]}
 */
function generatePaginationSchema(total, page, limit) {
    const totalPages = Math.ceil(total / limit);
    const paginationArray = [];

    if (totalPages <= 7) {
        for (let i = 1; i <= totalPages; i++) {
            paginationArray.push(i);
        }
    } else {
        paginationArray.push(1);

        if (page < 4) {
            for (let i = 2; i <= 4; i++) {
                paginationArray.push(i);
            }
            paginationArray.push(0);
        }

        if (page >= 4 && page <= totalPages - 3) {
            paginationArray.push(0);
            for (let i = page - 1; i <= page + 1; i++) {
                paginationArray.push(i);
            }
            paginationArray.push(0);
        }

        if (page > totalPages - 3) {
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
 * @param {number} total
 * @param {number} page
 * @param {number} limit
 * @returns {{
 * start: number,
 * end: number,
 * prev: number,
 * next: number,
 * total: number,
 * schema: number[]
 * }}
 */
export function pagination(total, page, limit) {
    const start = (page - 1) * limit + 1;
    const end = page * limit > total ? total : page * limit;
    const prev = page > 1 ? page - 1 : 1;
    const next =
        page < Math.ceil(total / limit) ? page + 1 : Math.ceil(total / limit);
    const schema = generatePaginationSchema(total, page, limit);

    return {
        start,
        end,
        prev,
        next,
        total,
        schema,
    };
}
