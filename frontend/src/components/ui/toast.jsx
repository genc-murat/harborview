import * as React from "react"
import * as ToastPrimitives from "@radix-ui/react-toast"
import { cn } from "@/lib/utils"
import PropTypes from "prop-types"

const ToastProvider = ToastPrimitives.Provider

const ToastViewport = React.forwardRef(({ className, ...props }, ref) => (
  <ToastPrimitives.Viewport
    ref={ref}
    className={cn(
      "fixed top-0 z-[100] flex max-h-screen w-full flex-col-reverse p-4 sm:bottom-0 sm:right-0 sm:top-auto sm:flex-col md:max-w-[420px]",
      className
    )}
    {...props}
  />
))

ToastViewport.propTypes = {
  className: PropTypes.string
}

ToastViewport.defaultProps = {
  className: ""
}

ToastViewport.displayName = ToastPrimitives.Viewport.displayName

const Toast = React.forwardRef(({ className, variant, children, ...props }, ref) => {
  return (
    <ToastPrimitives.Root
      ref={ref}
      className={cn(
        "group pointer-events-auto relative flex w-full items-center justify-between space-x-4 overflow-hidden rounded-md border p-6 pr-8 shadow-lg transition-all",
        {
          "border-gray-200 bg-white text-gray-950": variant === "default",
          "destructive group border-red-500 bg-red-600 text-gray-50": variant === "destructive",
        },
        className
      )}
      {...props}
    >
      {children}
    </ToastPrimitives.Root>
  )
})

Toast.propTypes = {
  className: PropTypes.string,
  variant: PropTypes.oneOf(["default", "destructive"]),
  children: PropTypes.node
}

Toast.defaultProps = {
  className: "",
  variant: "default"
}

Toast.displayName = ToastPrimitives.Root.displayName

const ToastAction = React.forwardRef(({ className, children, ...props }, ref) => (
  <ToastPrimitives.Action
    ref={ref}
    className={cn(
      "inline-flex h-8 shrink-0 items-center justify-center rounded-md border bg-transparent px-3 text-sm font-medium ring-offset-background transition-colors hover:bg-secondary focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 group-[.destructive]:border-muted/40 group-[.destructive]:hover:border-destructive/30 group-[.destructive]:hover:bg-destructive group-[.destructive]:hover:text-destructive-foreground group-[.destructive]:focus:ring-destructive",
      className
    )}
    {...props}
  >
    {children}
  </ToastPrimitives.Action>
))

ToastAction.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

ToastAction.defaultProps = {
  className: ""
}

ToastAction.displayName = ToastPrimitives.Action.displayName

const ToastClose = React.forwardRef(({ className, ...props }, ref) => (
  <ToastPrimitives.Close
    ref={ref}
    className={cn(
      "absolute right-2 top-2 rounded-md p-1 text-gray-950/50 opacity-0 transition-opacity hover:text-gray-950 focus:opacity-100 focus:outline-none focus:ring-2 group-hover:opacity-100 group-[.destructive]:text-red-300 group-[.destructive]:hover:text-red-50 group-[.destructive]:focus:ring-red-400 group-[.destructive]:focus:ring-offset-red-600",
      className
    )}
    toast-close=""
    {...props}
  />
))

ToastClose.propTypes = {
  className: PropTypes.string
}

ToastClose.defaultProps = {
  className: ""
}

ToastClose.displayName = ToastPrimitives.Close.displayName

const ToastTitle = React.forwardRef(({ className, children, ...props }, ref) => (
  <ToastPrimitives.Title
    ref={ref}
    className={cn("text-sm font-semibold", className)}
    {...props}
  >
    {children}
  </ToastPrimitives.Title>
))

ToastTitle.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

ToastTitle.defaultProps = {
  className: ""
}

ToastTitle.displayName = ToastPrimitives.Title.displayName

const ToastDescription = React.forwardRef(({ className, children, ...props }, ref) => (
  <ToastPrimitives.Description
    ref={ref}
    className={cn("text-sm opacity-90", className)}
    {...props}
  >
    {children}
  </ToastPrimitives.Description>
))

ToastDescription.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
}

ToastDescription.defaultProps = {
  className: ""
}

ToastDescription.displayName = ToastPrimitives.Description.displayName

export {
  ToastProvider,
  ToastViewport,
  Toast,
  ToastTitle,
  ToastDescription,
  ToastClose,
  ToastAction,
}