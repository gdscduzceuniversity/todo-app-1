const AuthLayout = ({ children }: { children: React.ReactNode }) => {
    return (
        <div className="flex items-center justify-center min-h-screen bg-white dark:bg-neutral-900 p-6">
            <div className="hidden bg-cover bg-center lg:block lg:min-w-[464px] h-[calc(100vh-48px)] rounded-lg"
                style={{
                    backgroundImage: `url(/assets/images/auth-layout-image.jpg)`
                }}
            />
            <div className="w-full flex justify-center">
                <div className="w-full sm:max-w-[464px]">
                    {children}
                </div>
            </div>
        </div>
    )
}

export default AuthLayout