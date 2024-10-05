"use client";
import * as z from 'zod';
import { useForm } from 'react-hook-form';
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

const formSchema = z.object({
  articleName: z.string().min(3, { message: "Article name must be at least 3 characters long" }),
  // Transform the input string into a number and validate
  price: z.preprocess((val) => parseFloat(val as string), z.number().nonnegative({ message: "Price must be a positive number" })),
    
  // Transform the input string into an integer and validate
  wears: z.preprocess((val) => parseInt(val as string, 10), z.number().int().nonnegative({ message: "Wears must be a non-negative integer" })),

  material: z.string(),
  brand: z.string().optional(),
  articleType: z.string(),
  season: z.string(),
  
})




export default function Home() {

  const articleTypeOptions = [
    { value: "top", label: "Top"},
    { value: "bottom", label: "Bottom"},
    { value: "shoes", label: "Shoes"},
    { value: "accessory", label: "accessory"},
  ]
  
  const seasonOptions = [
    { value: "summer", label: "Summer"},
    { value: "spring", label: "Spring"},
    { value: "fall", label: "Fall"},
    { value: "winter", label: "Winter"},
    { value: "yearRound", label: "Year Round"},
    { value: "cold", label: "Cold Weather"},
    { value: "warm", label: "Warm Weather"},
  ]

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      articleName: "",
      price: 0,
      wears: 0,
      brand: "",
      material: "",

    }
  })

  const handleSubmit = (data: z.infer<typeof formSchema>) => {
    console.log(data);
  }
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <h1>Add New Article</h1>
      <Form {...form}> 
        <form onSubmit={form.handleSubmit(handleSubmit)} className='max-w-md w-full'>
          <FormField control={form.control} name="articleName" render={({ field }) => (
              <FormItem>
                <FormLabel>Article Name</FormLabel>
                <FormControl>
                  <Input placeholder="Article Name" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )} />
          <div className='grid grid-cols-2 gap-2'>
            <FormField control={form.control} name="price" render={({ field }) => (
              <FormItem>
                <FormLabel>Price</FormLabel>
                <FormControl>
                  <Input type="number" step="0.01" placeholder="Price" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )} />

            <FormField control={form.control} name="wears" render={({ field }) => (
              <FormItem>
                <FormLabel>Wears</FormLabel>
                <FormControl>
                  <Input type="number" step="1" placeholder="Number of wears" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )} />

            <FormField control={form.control} name="brand" render={({ field }) => (
              <FormItem>
                <FormLabel>Brand Name</FormLabel>
                <FormControl>
                  <Input placeholder="Brand Name" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )} />

            <FormField control={form.control} name="material" render={({ field }) => (
              <FormItem>
                <FormLabel>Material</FormLabel>
                <FormControl>
                  <Input placeholder="Material" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )} />
            <FormField
              control={form.control}
              name="articleType"
              render={({ field }) => {
                return (
                  <FormItem>
                    <FormLabel>Article Type</FormLabel>
                    <Select onValueChange={field.onChange}>
                      <FormControl>
                        <SelectTrigger>
                          <SelectValue placeholder="Select an article type" />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        {articleTypeOptions.map((option) => (
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

          
          
          <Button type="submit" className='mt-4'>Submit</Button>
        </form>
      </Form> 
    </main>
  );
}
