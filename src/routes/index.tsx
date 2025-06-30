import { createFileRoute } from "@tanstack/react-router";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "~/components/ui/carousel";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "~/components/ui/card";
import { Button } from "~/components/ui/button";
import { Phone } from "lucide-react";

export const Route = createFileRoute("/")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <>
      {/* Hero Section with Carousel */}
      <section className="relative">
        <Carousel>
          <CarouselContent>
            <CarouselItem>
              <div className="relative w-full aspect-video overflow-hidden">
                <img
                  src="/placeholder.svg"
                  alt="Salon Interior"
                  className="absolute inset-0 w-full h-full object-cover"
                />
              </div>
            </CarouselItem>
            <CarouselItem>
              <div className="relative w-full aspect-video overflow-hidden">
                <img
                  src="/placeholder.svg"
                  alt="Stylish Haircut"
                  className="absolute inset-0 w-full h-full object-cover"
                />
              </div>
            </CarouselItem>
            <CarouselItem>
              <div className="relative w-full aspect-video overflow-hidden">
                <img
                  src="/placeholder.svg"
                  alt="Relaxing Spa"
                  className="absolute inset-0 w-full h-full object-cover"
                />
              </div>
            </CarouselItem>
          </CarouselContent>
          <CarouselPrevious className="absolute top-1/2 left-4 transform -translate-y-1/2">
            <Button variant="noShadow" size="icon">
              &lt;
            </Button>
          </CarouselPrevious>
          <CarouselNext className="absolute top-1/2 right-4 transform -translate-y-1/2">
            <Button variant="noShadow" size="icon">
              &gt;
            </Button>
          </CarouselNext>
        </Carousel>
      </section>

      {/* Services Section */}
      <section className="py-16 bg-secondary-background">
        <div className="container mx-auto grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
          <Card>
            <CardHeader>
              <CardTitle>Hair Styling</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                Expert haircuts and styling for all occasions.
              </CardDescription>
            </CardContent>
            <CardFooter>
              <Button variant="default" size="lg">
                Book Now
              </Button>
            </CardFooter>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Facial Treatments</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                Rejuvenating facials to refresh your skin.
              </CardDescription>
            </CardContent>
            <CardFooter>
              <Button variant="default" size="lg">
                Book Now
              </Button>
            </CardFooter>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Massage Therapy</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                Relaxing massages to relieve stress.
              </CardDescription>
            </CardContent>
            <CardFooter>
              <Button variant="default" size="lg">
                Book Now
              </Button>
            </CardFooter>
          </Card>
        </div>
      </section>

      {/* Contact Section */}
      <section className="py-16 bg-secondary-background">
        <div className="container mx-auto grid grid-cols-1 lg:grid-cols-2 gap-8">
          <div>
            <h2 className="text-3xl font-bold">Contact Us</h2>
            <p className="mt-4">
              Ready for your next beauty appointment? Call one of our contacts
              below to book your visit!
            </p>
            <ul className="mt-8 space-y-4">
              <li className="flex items-center">
                <Phone className="mr-2" />
                <span className="font-semibold">Thomas:</span>
                <span className="ml-2">+628123456789</span>
              </li>
              <li className="flex items-center">
                <Phone className="mr-2" />
                <span className="font-semibold">Sekar:</span>
                <span className="ml-2">+628164829372</span>
              </li>
            </ul>
          </div>
          <div>
            <h2 className="text-3xl font-bold">Location</h2>
            <p className="mt-4">Visit us at our salon located at:</p>
            <p className="mt-2">123 Beauty Street, Salon City, SC 12345</p>
            <div className="mt-8">
              <iframe
                src="https://www.google.com/maps/embed?pb=..."
                width="100%"
                height="400"
                style={{ border: 0 }}
                allowFullScreen
                loading="lazy"
              ></iframe>
            </div>
          </div>
        </div>
      </section>
    </>
  );
}
