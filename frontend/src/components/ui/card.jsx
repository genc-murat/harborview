import * as React from "react";
import PropTypes from "prop-types";
import { cn } from "@/lib/utils";

const Card = React.forwardRef(({ className, children, ...props }, ref) => (
  <div
    ref={ref}
    className={cn(
      "rounded-lg border bg-card shadow-sm transition-all hover:shadow-md",
      className
    )}
    {...props}
  >
    {children}
  </div>
));

Card.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

Card.defaultProps = {
  className: ""
}

Card.displayName = "Card"

const CardHeader = React.forwardRef(({ className, children, ...props }, ref) => (
  <div
    ref={ref}
    className={cn("flex flex-col space-y-1.5 p-6", className)}
    {...props}
  >
    {children}
  </div>
))

CardHeader.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

CardHeader.defaultProps = {
  className: ""
}

CardHeader.displayName = "CardHeader"

const CardTitle = React.forwardRef(({ className, children, ...props }, ref) => (
  <h3
    ref={ref}
    className={cn("text-lg font-semibold leading-none tracking-tight", className)}
    {...props}
  >
    {children}
  </h3>
))

CardTitle.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

CardTitle.defaultProps = {
  className: ""
}

CardTitle.displayName = "CardTitle"

const CardDescription = React.forwardRef(({ className, children, ...props }, ref) => (
  <p
    ref={ref}
    className={cn("text-sm text-muted-foreground", className)}
    {...props}
  >
    {children}
  </p>
))

CardDescription.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

CardDescription.defaultProps = {
  className: ""
}

CardDescription.displayName = "CardDescription"

const CardContent = React.forwardRef(({ className, children, ...props }, ref) => (
  <div 
    ref={ref} 
    className={cn("p-6 pt-0", className)} 
    {...props}
  >
    {children}
  </div>
))

CardContent.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

CardContent.defaultProps = {
  className: ""
}

CardContent.displayName = "CardContent"

const CardFooter = React.forwardRef(({ className, children, ...props }, ref) => (
  <div
    ref={ref}
    className={cn("flex items-center p-6 pt-0", className)}
    {...props}
  >
    {children}
  </div>
))

CardFooter.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

CardFooter.defaultProps = {
  className: ""
}

CardFooter.displayName = "CardFooter"

export { Card, CardHeader, CardFooter, CardTitle, CardDescription, CardContent }