import Image from "next/image";
export default function Wallpaper() {
  return (
    <div className="absolute inset-0 w-full h-full">
    <Image
    src={"/wallpapers/dark_mountains.jpg"}
    alt=""
    fill
    draggable={false}
    className="blur-sm opacity-80"
    />
    </div>
  );
}
