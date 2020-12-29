import React, { ReactElement, ReactNode, useState } from 'react';

export type TableRowProps = {
    type: 'alert' | null;
    row: {
        getRowProps: () => {
            key: string;
        };
        isGrouped: boolean;
    };
    children: ReactNode;
    HoveredRowComponent?: ReactNode;
    GroupedRowComponent?: ReactNode;
};

const tableRowClassName = 'relative border-b';
const baseTableRowClassName = `${tableRowClassName} border-base-300`;
const alertTableRowClassName = `${tableRowClassName} border-alert-300 bg-alert-200 text-alert-800`;

function onFocus(): number {
    return 0;
}

function TableRowOverlay({ children }): ReactElement {
    return (
        <td className="absolute right-0 transform -translate-x-2 translate-y-1 mr-2">{children}</td>
    );
}

function TableRow({
    type,
    row,
    children,
    HoveredRowComponent = null,
    GroupedRowComponent = null,
}: TableRowProps): ReactElement {
    const [isHovered, setIsHovered] = useState(false);

    const { key } = row.getRowProps();
    const className = type === 'alert' ? alertTableRowClassName : baseTableRowClassName;
    const showGroupedRowComponent = row.isGrouped;
    const showHoveredRowComponent = !showGroupedRowComponent && isHovered;

    const hoveredRowComponent = showHoveredRowComponent && (
        <TableRowOverlay>{HoveredRowComponent}</TableRowOverlay>
    );

    const groupedRowComponent = showGroupedRowComponent && (
        <TableRowOverlay>{GroupedRowComponent}</TableRowOverlay>
    );

    function onMouseEnter(): void {
        setIsHovered(true);
    }

    function onMouseLeave(): void {
        setIsHovered(false);
    }

    return (
        <tr
            key={key}
            className={className}
            onMouseEnter={onMouseEnter}
            onMouseLeave={onMouseLeave}
            onFocus={onFocus}
        >
            {children}
            {hoveredRowComponent}
            {groupedRowComponent}
        </tr>
    );
}

export default TableRow;
