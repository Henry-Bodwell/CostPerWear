
import React from 'react';
import {Button} from "@/components/ui/button";

interface WearButtonProps {
    id: number;
    onWear: () => void;
}

const WearButton: React.FC<WearButtonProps> = ({id, onWear}) => {
    const handleWear = async () => {
        try {
            const response = await fetch(`http://localhost:9090/api/clothes/${id}`, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json'
                },
            });

            if (!response.ok) {
                throw new Error('Failed to wear article');
            }

            const data = await response.json();
            console.log('Wear updated:', data);
            onWear();
        } catch (error) {
            console.error('Error updating wear:', error);
        }

    };

    

    return (
        <Button onClick={handleWear}>Wear</Button>
    );
}

export default WearButton;