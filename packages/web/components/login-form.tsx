"use client"

import { Button } from "@/components/ui/button"
import { Field, FieldLabel } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { HugeiconsIcon } from "@hugeicons/react"
import { SakuraIcon } from "@hugeicons/core-free-icons"

export function LoginForm({ className, ...props }: React.ComponentProps<"div">) {
  return (
    <div className={`flex flex-col items-center ${className ?? ""}`} {...props}>
      <div className="flex flex-col items-center mb-4">
        <HugeiconsIcon icon={SakuraIcon} className="h-10 w-10 mb-2" />
        <h1 className="text-2xl font-bold">Create Profile</h1>
      </div>
      <form className="w-full max-w-sm flex flex-col gap-4">
        <Field>
          <FieldLabel htmlFor="username">Username</FieldLabel>
          <Input id="username" type="text" placeholder="Username" required />
        </Field>
        <Field>
          <FieldLabel htmlFor="password">Password</FieldLabel>
          <Input id="password" type="password" placeholder="•••••••••" required />
        </Field>
        <Field>
          <FieldLabel htmlFor="confirm_password">Confirm Password</FieldLabel>
          <Input id="confirm_password" type="password" placeholder="•••••••••" required />
        </Field>
        <span className="text-xs text-muted-foreground text-start">
          Write down your password. This can never be recovered.
        </span>
        <Button type="submit" className="w-full">Login</Button>
      </form>
    </div>
  )
}
