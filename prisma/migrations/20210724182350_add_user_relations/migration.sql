-- AlterTable
ALTER TABLE "EducationItem" ADD COLUMN     "userId" TEXT;

-- AlterTable
ALTER TABLE "List" ADD COLUMN     "userId" TEXT;

-- AddForeignKey
ALTER TABLE "EducationItem" ADD FOREIGN KEY ("userId") REFERENCES "User"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "List" ADD FOREIGN KEY ("userId") REFERENCES "User"("id") ON DELETE SET NULL ON UPDATE CASCADE;
