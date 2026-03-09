import { LoginForm } from "@/components/login-form";
import Wallpaper from "@/components/ui/wallpaper";

export default function Home() {
  return (
    <div className="flex min-h-svh flex-col items-center justify-center gap-6 bg-background p-6 md:p-10">
      <Wallpaper />
      <div className="w-full max-w-sm z-9">
        <LoginForm />
      </div>
    </div>
  );
}
