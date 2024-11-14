"use client";

import ArticleCards from "@/components/ArticleCards"

import Link from "next/link";
import { NavigationMenu, NavigationMenuItem, NavigationMenuLink, navigationMenuTriggerStyle } from "@/components/ui/navigation-menu";


export default function main() {
  return (
    <div>
      <header>
        <NavigationMenu>
          <NavigationMenuItem>
            <Link href="/" legacyBehavior passHref>
              <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                Home
              </NavigationMenuLink>
            </Link>
          </NavigationMenuItem>
          <NavigationMenuItem>
            <Link href="/submit" legacyBehavior passHref>
              <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                New Item
              </NavigationMenuLink>
            </Link>
          </NavigationMenuItem>
        </NavigationMenu>
      </header>
      <main>
        <ArticleCards />
      </main>
    </div>
  )
}