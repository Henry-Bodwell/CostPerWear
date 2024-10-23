"use client";
import * as z from 'zod';
import { SubmitHandler, useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import {
  SelectValue,
  SelectTrigger,
  SelectContent,
  SelectItem,
  Select,
} from "@/components/ui/select";


const OutfitForm = () => {
  const formSchema = z.object({
    outfitName: z.string().min(3),

    vibes: z.string().optional(),
    season: z.string(),
  })

  const seasonOptions = [
    { value: "summer", label: "Summer"},
    { value: "spring", label: "Spring"},
    { value: "fall", label: "Fall"},
    { value: "winter", label: "Winter"},
    { value: "yearRound", label: "Year Round"},
    { value: "cold", label: "Cold Weather"},
    { value: "warm", label: "Warm Weather"},
  ] 

  const vibeOptions = [
    { value: "formal", label: "Formal"},
    { value: "casual", label: "Casual"},
    { value: "comfy", label: "Comfy"},
  ]

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      outfitName: "",
      vibes: "",
      season: "",
    }
  })

  const {isSubmitting} = form.formState;

  const handleSubmit: SubmitHandler<z.infer<typeof formSchema>> = async (data: z.infer<typeof formSchema>) => {
    await new Promise ((resolve) => setTimeout(resolve, 1000))
    console.log(data);
  }

  return (
    <main className="flex min-h-screen flex-col items-center py-20">
    <h1>Add New Article</h1>
    <Form {...form}> 
        <form onSubmit={form.handleSubmit(handleSubmit)} className='max-w-md w-full'>
        <FormField
         control={form.control} 
         name="outfitName" 
         render={({ field }) => (
            <FormItem>
                <FormLabel>Article Name</FormLabel>
                <FormControl>
                <Input placeholder="Outfit Name" {...field} />
                </FormControl>
                <FormMessage />
            </FormItem>
            )} />
        <div className='grid grid-cols-2 gap-2'>
            <FormField
            control={form.control}
            name="vibes"
            render={({ field }) => {
                return (
                <FormItem>
                    <FormLabel>Article Type</FormLabel>
                    <Select onValueChange={field.onChange}>
                    <FormControl>
                        <SelectTrigger>
                        <SelectValue placeholder="Select a vibe" />
                        </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                        {vibeOptions.map((option) => (
                        <SelectItem key={option.value} value={option.value}>
                        {option.label}
                        </SelectItem>
                    ))}
                    </SelectContent>
                    </Select>
                    <FormMessage />
                </FormItem>
                );
            }} />
            <FormField
            control={form.control}
            name="season"
            render={({ field }) => {
                return (
                <FormItem>
                    <FormLabel>Season</FormLabel>
                    <Select onValueChange={field.onChange}>
                    <FormControl>
                        <SelectTrigger>
                        <SelectValue placeholder="Select a season" />
                        </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                        {seasonOptions.map((option) => (
                        <SelectItem key={option.value} value={option.value}>
                        {option.label}
                        </SelectItem>
                    ))}
                    </SelectContent>
                    </Select>
                    <FormMessage />
                </FormItem>
                );
            }} />
        </div>
        <Button disabled={isSubmitting} type="submit" className='mt-4'>
            {isSubmitting ? "Loading..." : "Submit"}
        </Button>
        </form>
    </Form> 
    </main>
    );
  }

export default OutfitForm;