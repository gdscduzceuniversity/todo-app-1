import AuthLayout from '@/src/views/pages/auth/auth-layout'
import { zodResolver } from '@hookform/resolvers/zod'
import { SubmitHandler, useForm } from 'react-hook-form'
import * as z from "zod"

import { Button } from "@/src/components/ui/button"
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/src/components/ui/form"
import { Input } from "@/src/components/ui/input"
import Link from 'next/link'
import { Checkbox } from '@/src/components/ui/checkbox'

const loginSchema = z.object({
    email: z.string().email({
        message: "Invalid email address"
    }),
    password: z.string().min(8, {
        message: "Password must be at least 8 characters"
    }),
    remember: z.boolean().optional()
})

type LoginSchema = z.infer<typeof loginSchema>

const LoginPage = () => {
    const form = useForm<LoginSchema>({
        resolver: zodResolver(loginSchema),
        defaultValues: {
            email: "",
            password: "",
            remember: false
        }
    })

    const onSubmit = (data: LoginSchema) => {
        console.log(data)
    }

    return (
        <AuthLayout>
            <div className="flex flex-col justify-center">
                <h1 className="mb-2">👋 Welcome back</h1>
                <p className="mb-4">
                    Please login to your account.
                </p>
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-4'>
                        <FormField
                            control={form.control}
                            name="email"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Email</FormLabel>
                                    <FormControl>
                                        <Input placeholder="Enter your email" {...field} />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="password"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Password</FormLabel>
                                    <FormControl>
                                        <Input type='password' placeholder="Enter your password" {...field} />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <div className='flex justify-between'>
                            <FormField control={form.control} name="remember" render={({ field }) => (
                                <FormItem className="flex justify-center items-center space-x-2">
                                    <FormControl>
                                        <Checkbox
                                            checked={field.value}
                                            onCheckedChange={field.onChange}
                                        />
                                    </FormControl>
                                    <FormLabel className='text-primary font-normal !mt-0'>
                                        Remember me
                                    </FormLabel>
                                </FormItem>
                            )} />
                            <Link href="/auth/forgot-password" className="text-muted hover:underline text-sm">
                                Forgot password?
                            </Link>
                        </div>
                        <Button type="submit" size="lg" className='w-full'>
                            Login
                        </Button>
                    </form>
                </Form>
                <div className="flex items-center justify-center space-x-2 my-4">
                    <p className="text-sm">Don't have an account?</p>
                    <Link href="/auth/register" className="text-primary hover:underline text-sm">
                        Sign up
                    </Link>
                </div>
            </div>
        </AuthLayout>
    )
}

export default LoginPage