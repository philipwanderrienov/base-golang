import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"


/*************  ✨ Windsurf Command ⭐  *************/
/**
 * A utility function that merges multiple class names into a single string.
 * It's a wrapper around `clsx` and `tailwind-merge`'s `twMerge` function.
 * It's useful for creating dynamic class names based on multiple conditionals.
 * @param {ClassValue[]} inputs - A variable number of class names to merge.
 * @returns {string} - A single string containing all the merged class names.
 */
/*******  6a438506-3a9c-4369-860f-b343a7f77e85  *******/export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}