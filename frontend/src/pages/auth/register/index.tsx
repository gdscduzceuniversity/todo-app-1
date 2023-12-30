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

const registerSchema = z.object({
    email: z.string().email({
        message: "Invalid email address"
    }),
    password: z.string().min(8, {
        message: "Password must be at least 8 characters"
    }),
    password_confirm: z.string().min(8, {
        message: "Password must be at least 8 characters"
    }),
    terms: z.boolean().refine((data) => data === true, {
        message: "You must agree to the terms and conditions",
    })
}).superRefine(({ password_confirm, password }, ctx) => {
    if (password_confirm !== password) {
        ctx.addIssue({
            code: "custom",
            message: "Passwords do not match",
        });
    }
});

type RegisterSchema = z.infer<typeof registerSchema>

const LoginPage = () => {
    const form = useForm<RegisterSchema>({
        resolver: zodResolver(registerSchema),
        defaultValues: {
            email: "",
            password: "",
            password_confirm: "",
            terms: false
        }
    })

    const onSubmit = (data: RegisterSchema) => {
        console.log(data)
    }

    return (
        <AuthLayout>
            <div className="flex flex-col justify-center">
                <h1 className="mb-2">👋 Join the App</h1>
                <p className="mb-4">
                    Fill in the form to get started
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
                        <FormField
                            control={form.control}
                            name="password_confirm"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Confirm Password</FormLabel>
                                    <FormControl>
                                        <Input type='password' placeholder="Enter your password again" {...field} />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <div className='flex justify-between'>
                            <FormField control={form.control} name="terms" render={({ field }) => (
                                <FormItem className="flex justify-center items-center space-x-2">
                                    <FormControl>
                                        <Checkbox
                                            checked={field.value}
                                            onCheckedChange={field.onChange}
                                        />
                                    </FormControl>
                                    <FormLabel className='text-primary font-normal !mt-0'>
                                        I agree to the <Link href="#" className="text-primary hover:underline text-sm">
                                            Terms of Service
                                        </Link>
                                    </FormLabel>
                                </FormItem>
                            )} />
                        </div>
                        <Button type="submit" size="lg" className='w-full'>
                            Sign Up
                        </Button>
                    </form>
                </Form>
                <div className="flex items-center justify-center space-x-2 my-4">
                    <p className="text-sm">Already have an account?</p>
                    <Link href="/auth/login" className="text-primary hover:underline text-sm">
                        Sign In
                    </Link>
                </div>
            </div>
        </AuthLayout>
    )
}

export default LoginPage