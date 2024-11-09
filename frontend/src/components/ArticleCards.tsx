import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import React, { useEffect, useState } from "react";
import WearButton from "./WearButton";
import { Button } from "@/components/ui/button";


interface ClothingArticle {
    name: string;
    price: number;
    wears: number;
    costPerWear: number;
    articleType: string;
    imageUrl: string;
    id: number;
    lastWorn: string;
}

const ArticleCards: React.FC = () => {
    const [data, setData] = useState<ClothingArticle[]>([]);

    const fetchArticles = async () => {
        try {
            const response = await fetch("http://localhost:9090/api/clothes");
            if (!response.ok) {
                throw new Error('Failed to fetch articles');
            }
            const data = await response.json();
            setData(data);
        } catch (error) {
            console.error('Error fetching articles', error);
        }
    };

    useEffect(() => {
        fetchArticles();
    }, []);

    const handleWear = async (id: number) => {
        setData(data.map(article => {
            if (article.id === id) {
                return {
                    ...article,
                    wears: article.wears + 1,
                    costPerWear: article.price / (article.wears + 1),
                };
            }
        return article;
        }));
    };
    return (

        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 p-4">
            {data.map((article, index) => (
                <Card key={index} className="bg-white rounded-lg shadow-md overflow-hidden">
                    <CardHeader>
                        <CardTitle className="grid grid-cols-2 gap-8">{article.name} <Button>View Article</Button></CardTitle>
                        <CardDescription>{article.costPerWear}</CardDescription>
                    </CardHeader>
                <CardContent>
                    <p>Price: {article.price} Wears: {article.wears} </p>
                    <p>Last Worn: {article.lastWorn}</p>
                </CardContent>
                <CardFooter className="grid grid-cols-2 gap-8">
                    <WearButton id={article.id} onWear={() => handleWear(article.id)}/>
                    <p>Article Type: {article.articleType}</p>
                </CardFooter>
            </Card>
            ))}
        </div>
        
    );
};

  


export default ArticleCards;