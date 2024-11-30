import * as React from "react"
import PropTypes from "prop-types"
import { cn } from "@/lib/utils"

const Table = React.forwardRef(({ className, children, ...props }, ref) => (
  <div className="relative w-full overflow-auto">
    <table
      ref={ref}
      className={cn("w-full caption-bottom text-sm", className)}
      {...props}
    >
      {children}
    </table>
  </div>
))

Table.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

Table.defaultProps = {
  className: ""
}

Table.displayName = "Table"

const TableHeader = React.forwardRef(({ className, children, ...props }, ref) => (
  <thead ref={ref} className={cn("[&_tr]:border-b", className)} {...props}>
    {children}
  </thead>
))

TableHeader.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

TableHeader.defaultProps = {
  className: ""
}

TableHeader.displayName = "TableHeader"

const TableBody = React.forwardRef(({ className, children, ...props }, ref) => (
  <tbody
    ref={ref}
    className={cn("[&_tr:last-child]:border-0", className)}
    {...props}
  >
    {children}
  </tbody>
))

TableBody.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

TableBody.defaultProps = {
  className: ""
}

TableBody.displayName = "TableBody"

const TableRow = React.forwardRef(({ className, children, ...props }, ref) => (
  <tr
    ref={ref}
    className={cn(
      "border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted",
      className
    )}
    {...props}
  >
    {children}
  </tr>
))

TableRow.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

TableRow.defaultProps = {
  className: ""
}

TableRow.displayName = "TableRow"

const TableHead = React.forwardRef(({ className, children, ...props }, ref) => (
  <th
    ref={ref}
    className={cn(
      "h-12 px-4 text-left align-middle font-medium text-muted-foreground [&:has([role=checkbox])]:pr-0",
      className
    )}
    {...props}
  >
    {children}
  </th>
))

TableHead.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

TableHead.defaultProps = {
  className: ""
}

TableHead.displayName = "TableHead"

const TableCell = React.forwardRef(({ className, children, ...props }, ref) => (
  <td
    ref={ref}
    className={cn("p-4 align-middle [&:has([role=checkbox])]:pr-0", className)}
    {...props}
  >
    {children}
  </td>
))

TableCell.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

TableCell.defaultProps = {
  className: ""
}

TableCell.displayName = "TableCell"

export {
  Table,
  TableHeader,
  TableBody,
  TableHead,
  TableRow,
  TableCell,
}